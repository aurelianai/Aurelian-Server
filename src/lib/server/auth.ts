import type { RequestEvent } from '@sveltejs/kit'
import { prisma } from '$lib/server/prisma-client'
import jwt from 'jsonwebtoken'


export class AuthStore {
   user_id: number | null = null
   token = ''
   isValid = false
   JWT_SECRET = process.env.JWT_SECRET || ''

   async load_from_cookie(event: RequestEvent): Promise<void> {
      const cookie = event.cookies.get("Authorization")
      if (cookie) {
         console.log(`Got Cookie ${cookie}`)
         this.token = cookie // Dump Bearer prefix
         console.log(`Got Token: '${this.token}`)
         try {
            const jwt_json = jwt.verify(this.token, this.JWT_SECRET) as JWT_BODY
            this.user_id = jwt_json.user_id
            this.isValid = true
         } catch (_) {
            console.log(`Token Verification Failed`)
         }
      }
   }

   refresh(): boolean {
      if (this.token && this.user_id && this.isValid) {
         this.token = jwt.sign({ user_id: this.user_id }, this.JWT_SECRET, { expiresIn: "1d" })
         return true
      }
      return false
   }

   export_to_cookie(): string {
      return `Authorization=${this.token}`
   }

   async hydrate_with_email_password(email: string, password: string): Promise<boolean> {
      const user = await prisma.user.findUnique({
         where: {
            email
         }
      })
      if (!user) {
         console.log("user not found")
         return false
      }

      if (password !== user.password) {
         console.log(`password didn't match ${user.password}`)
         return false
      }

      this.user_id = user.id
      this.token = jwt.sign({ user_id: this.user_id }, this.JWT_SECRET, { expiresIn: "1d" })
      this.isValid = true
      return true
   }

   clear(): void {
      this.user_id = null
      this.token = ''
      this.isValid = false
   }
}

type JWT_BODY = {
   user_id: number
} 
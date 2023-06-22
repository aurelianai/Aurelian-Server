import type { RequestEvent } from '@sveltejs/kit'
import { prisma } from '$lib/server/prisma-client'
import jwt from 'jsonwebtoken'
import bcrypt from 'bcrypt'


export class AuthStore {
   user_id: number | null = null
   token: string = ''
   isValid: boolean = false
   JWT_SECRET = process.env.JWT_SECRET || ''

   load_from_cookie(event: RequestEvent): void {
      let cookie = event.cookies.get("Authorization")
      if (cookie) {
         this.token = cookie.split(' ')[0] // Dump Bearer prefix
         const decoded_token_body = jwt.verify(this.token, this.JWT_SECRET) as JWT_BODY
         this.user_id = decoded_token_body.user_id
         this.isValid = true
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
      return `Authorization=Bearer ${this.token}`
   }

   async hydrate_with_email_password(email: string, password: string): Promise<boolean> {
      const user = await prisma.user.findUnique({
         where: {
            email
         }
      })
      if (!user) {
         return false
      }

      const valid = await bcrypt.compare(password, user.password)
      if (!valid) {
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
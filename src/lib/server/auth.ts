import type { RequestEvent } from '@sveltejs/kit'
import jwt from 'jsonwebtoken'


export class AuthStore {
   user_id: number | null = null
   token: string | null = null
   isValid: boolean = false

   load_from_cookie(event: RequestEvent): void {
      let cookie = event.cookies.get("JWT_TOKEN")
      if (cookie) {
         token = cookie.split(" ")[0]
         try {
            user_id = jwt.verify(token, process.env.JWT_SECRET).user_id
            isValid = true
         }
      }
   }

   refresh(): boolean {
      if (token && user_id && isValid) {
         token = jwt.sign({ user_id: user_id }, process.env.JWT_SECRET, { epxiresIn: "1d" })
         return true
      }
      return false
   }

   export_to_cookie(event: RequestEvent): void {
      event.cookies.set("JWT_TOKEN", token)
   }
}
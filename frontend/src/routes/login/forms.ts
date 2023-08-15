import { z } from 'zod'

export const validate_login = (email: string, pass: string): string | null => {
   const form_schema = z.object({
      email: z.string().min(5).email("Email must be valid"),
      pass: z.string().min(5),
   })
   try {
      form_schema.parse({ email: email, pass: pass })
   } catch (e: any) {
      return e.issues.map((el: any) => el.message).join('\n')
   }
   return null
}
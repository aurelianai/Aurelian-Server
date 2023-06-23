import type { ServerLoad, Actions } from '@sveltejs/kit'
import type { ZodError } from 'zod'
import type { LoginForm } from './types'
import { redirect } from '@sveltejs/kit'
import { z } from 'zod'

export const load = (async ({ locals }) => {
   if (locals.auth_store.isValid) {
      throw redirect(307, '/chat')
   }
}) satisfies ServerLoad


const loginSchema = z.object({
   email: z.string({ required_error: "Please enter an email" }).email({ message: "Please enter a valid email" }),
   password: z.string({ required_error: "Please enter a password" }).trim(),
})



export const actions: Actions = {
   login: async ({ request, locals }) => {
      const form = Object.fromEntries(await request.formData()) as LoginForm
      try {
         loginSchema.parse(form)
      } catch (e) {
         const { fieldErrors: errors } = (e as ZodError).flatten()
         return {
            data: { email: form.email },
            errors: errors
         }
      }

      if (await locals.auth_store.hydrate_with_email_password(form.email, form.password)) {
         throw redirect(303, "/chat")
      }
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { password, ...rest } = form
      return {
         data: rest,
         errors: {
            email: ["Invalid Email or Password"],
            password: ["Invalid Email or Password"]
         }
      }
   }
}
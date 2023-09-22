import type { PageServerLoad, Actions } from './$types'
import { fail, redirect } from '@sveltejs/kit'
import { LuciaError } from 'lucia'
import { z } from 'zod'
import { auth } from '$lib/server/auth'

export const load: PageServerLoad = async ({ url, locals }) => {
   const session = await locals.auth.validate();
   if (session) throw redirect(302, "/chat")

   return {
      'ref': url.searchParams.get("ref"),
      'username': url.searchParams.get("username"),
      'password': url.searchParams.get("password"),
   }
}

export const actions: Actions = {
   login: async ({ request, locals }) => {
      const formData = await request.formData()
      const fusername = formData.get("username")
      const fpassword = formData.get("password")

      const loginValidator = z.object({
         username: z.string().min(5).max(20),
         password: z.string().min(5),
      })

      const loginValidation = loginValidator.safeParse({ username: fusername, password: fpassword })

      if (!loginValidation.success) {
         const errors = loginValidation.error.errors.map((e) => {
            return {
               field: e.path[0],
               message: e.message
            }
         })
         return fail(400, { username: fusername, errors })
      }

      const username = loginValidation.data.username
      const password = loginValidation.data.password
      try {
			const key = await auth.useKey(
				"username",
				username.toLowerCase(),
				password
			);
			const session = await auth.createSession({
				userId: key.userId,
				attributes: {}
			});
			locals.auth.setSession(session);
		} catch (e) {
			if (
				e instanceof LuciaError &&
				(e.message === "AUTH_INVALID_KEY_ID" ||
					e.message === "AUTH_INVALID_PASSWORD")
         ) {
            return fail(400, {
               username,
               errors: [{
                  field: "username",
                  message: "Incorrect Username or Password"
               }]
            });
			}
         return fail(500, {
            username,
            errors: [{
               field: "500",
               message: "An unkown error occured"
            }]
			});
		}
		throw redirect(302, "/chat");
   }
}
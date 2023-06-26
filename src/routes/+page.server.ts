import type { ServerLoad } from '@sveltejs/kit'
import { redirect } from '@sveltejs/kit'

export const load = (async ({ locals }) => {
   if (!locals.auth_store.isValid) {
      throw redirect(307, "/login")
   } else {
      throw redirect(303, '/chat')
   }
}) satisfies ServerLoad
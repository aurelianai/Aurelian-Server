import type { ServerLoad } from '@sveltejs/kit'
import { redirect } from '@sveltejs/kit'

export const load = (async ({ locals }) => {
   locals.auth_store.clear()
   throw redirect(303, '/login')
}) satisfies ServerLoad
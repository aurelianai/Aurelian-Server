import { redirect, json, type RequestHandler } from '@sveltejs/kit'
import { supported_models } from '$lib/server/completions/models'

export const GET = (async ({ locals }) => {
   if (!locals.auth_store.isValid) {
      throw redirect(307, "/login")
   }
   return json(supported_models)
}) satisfies RequestHandler
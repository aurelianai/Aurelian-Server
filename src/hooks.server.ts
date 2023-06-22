// Start up code to migrate database. Not the cleanest, but should get the job done.
import { execSync } from 'child_process'
execSync(`npx prisma migrate deploy --schema /aels/prisma/schema.prisma`, { stdio: "inherit" })


// Auth Middleware
import type { Handle } from '@sveltejs/kit'
import { AuthStore } from '$lib/server/auth'


export const handle = (async ({ event, resolve }) => {
   event.locals.auth_store = new AuthStore()
   event.locals.auth_store.load_from_cookie(event)

   if (event.locals.auth_store.isValid) {
      event.locals.auth_store.refresh()
   } else {
      event.locals.auth_store.clear()
   }

   const response = await resolve(event)

   response.headers.append('Set-Cookie', event.locals.auth_store.export_to_cookie())

   return response
}) satisfies Handle
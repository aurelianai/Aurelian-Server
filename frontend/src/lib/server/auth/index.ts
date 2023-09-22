import { lucia } from 'lucia'
import { sveltekit } from 'lucia/middleware'
import { pg } from '@lucia-auth/adapter-postgresql'
import { adb } from '$lib/server/persistence'
import { dev } from '$app/environment'

export const auth = lucia({
   env: dev ? "DEV" : "PROD",
   middleware: sveltekit(),
   adapter: pg(adb.pool, {
      user: 'user',
      key: 'key',
      session: 'session'
   }),

   getUserAttributes: (data) => {
      return { username: data.username }
   }
})

export type Auth = typeof auth

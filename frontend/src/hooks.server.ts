/* Server Start Up Code */
import { pg } from '@lucia-auth/adapter-postgresql'
import { lucia } from 'lucia'
import { adb } from '$lib/server/persistence'

await adb.migrateToLatest()

const auth = lucia({
   adapter: pg(adb.pool, {
      user: "users",
      key: "user_key",
      session: "user_session"
   }),
   env: "DEV"
})



/* Server Hooks */
import type { Handle } from '@sveltejs/kit'

export const handle: Handle = async ({ event, resolve }) => {
   return resolve(event)
}
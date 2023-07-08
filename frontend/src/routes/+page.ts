import type { PageLoad } from './$types'
import { goto } from '$app/navigation'

export const load = (async () => {
   let res = await fetch("/api/auth")
   if (res.status == 200) {
      goto("/chat")
   } else {
      goto("/login")
   }
}) satisfies PageLoad
import type { PageLoad } from './$types'
import { goto } from '$app/navigation'

export const load = (async({ fetch }) => {
   let res = await fetch("/api/auth")
   if (res.status == 200) {
      await goto("/chat")
   } else {
      await goto("/login")
   }
}) satisfies PageLoad
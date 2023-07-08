import { goto } from '$app/navigation'
import type { PageLoad } from './$types'

export const load = (async ({ url }) => {
   let res = await fetch("/api/auth")
   if (res.status === 200) {
      goto("/chat")
   }

   return {
      'ref': url.searchParams.get("ref"),
      'email': url.searchParams.get("email"),
      'pass': url.searchParams.get("pass"),
   }
}) satisfies PageLoad
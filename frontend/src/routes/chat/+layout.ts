import type { LayoutLoad } from './$types'
import type { Chat } from '$lib/types'
import { goto } from '$app/navigation'
import { ChatStore } from '$lib/ts/chat/util'

export const load = (async ({ fetch }) => {
   let res = await fetch("/api/chat")
   if (res.status == 401) {
      goto("/login")
   }
   const chats: Chat[] = await res.json()
   ChatStore.set(chats)
   return { chats }
}) satisfies LayoutLoad
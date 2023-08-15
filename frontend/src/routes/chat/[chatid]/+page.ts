import type { PageLoad } from './$types'
import type { Message } from '$lib/types'
import { goto } from '$app/navigation'
import { error } from '@sveltejs/kit'
import { get } from 'svelte/store'
import { ChatStore } from '$lib/ts/chat/util'


export const load = (async ({ fetch, params }) => {
   if (isNaN(+params.chatid)) {
      throw error(400)
   }
   let res = await fetch(`/api/chat/${params.chatid}`)
   if (res.status == 401) {
      goto('/login')
   }

   const chatTitle = get(ChatStore)?.find((c) => c.ID === +params.chatid)?.Title

   let messages: Message[] = await res.json()
   return { chatid: +params.chatid, chatTitle, messages }
}) satisfies PageLoad
import type { Chat, Message, User } from '$lib/types'
import { writable } from 'svelte/store'


export const new_chat = async (title: string): Promise<Chat> => {
   const res = await fetch("/api/chat", {
      headers: new Headers({ 'content-type': 'application/json' }),
      method: "POST",
      body: JSON.stringify({ title })
   })
   return res.json()
}

export const update_chat = async (chatid: number, new_title: string) => {
   const body = { new_title }
   await fetch(`/api/chat?chatid=${chatid}`, {
      headers: new Headers({ 'content-type': 'application/json' }),
      method: "PATCH",
      body: JSON.stringify(body)
   })
}

export const delete_chat = async (id: number) => {
   const res = await fetch(`/api/chat?chatid=${id}`, {
      headers: new Headers({ 'content-type': 'application/json' }),
      method: "DELETE"
   })
}

export const new_message = async (chatid: number, role: "USER" | "MODEL", content: string): Promise<Message> => {
   const body: Message = { Role: role, Content: content }
   const res = await fetch(`/api/chat/${chatid}`, {
      headers: new Headers({ 'content-type': 'application/json' }),
      method: "POST",
      body: JSON.stringify(body)
   })
   return res.json()
}

type StreamResponse = {
	token: Token
}

type Token = {
	id: number    
	text: string 
	logprob: number 
	special: boolean 
}

export async function* complete(chatid: number, sig: AbortSignal): AsyncGenerator<string> {

   const response = await fetch(`/api/chat/${chatid}/complete`, {
      headers: new Headers({ 'content-type': 'application/json' }),
      method: "POST" 
   })

   if (!response.body) {
      throw new Error("No Body attached to response")
   }

   let buffer = ''
   const reader = response.body.getReader()
   const decoder = new TextDecoder('utf-8')

   while (true) {
      if (sig.aborted) {
         return
      }

      const { done, value } = await reader.read()
      if (done) { return }

      buffer += decoder.decode(value)

      while (buffer.includes("\n\n")) {
         const lineEnd = buffer.indexOf("\n\n")
         const rawJson = buffer.slice(5, lineEnd)
         buffer = buffer.slice(lineEnd + 2)

         const streamResponse: StreamResponse = JSON.parse(rawJson)
         yield streamResponse.token.text
      }
   }

}

export const ChatStore = writable<Chat[]>()
export const UserStore = writable<User>()

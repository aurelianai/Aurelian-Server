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

export type InferenceUpdate = {
   delta: string,
   err: string,
   last: boolean,
}

export async function* complete(chatid: number, sig: AbortSignal): AsyncGenerator<InferenceUpdate> {

   const response = await fetch(`/api/chat/${chatid}/complete`, {
      headers: new Headers({ 'content-type': 'application/json' }),
      method: "POST",
      signal: sig
   })

   if (!response.body) {
      throw new Error("No Body attached to response")
   }

   for await (let line of iterLines(response.body)) {
      try {
         const update = JSON.parse(line)
         yield update 
      } catch (e) {
         yield {
            delta: "", err: `error occured during parsing of line: '${line}'. Error: ${JSON.stringify(e)}`, last: false
         }
      }
   }
}


export async function* iterLines(stream: ReadableStream<Uint8Array>): AsyncGenerator<string> {
   let pending: string | undefined = undefined
   const reader = stream.getReader()
   const decoder = new TextDecoder('utf-8')

   while (true) {

      const { done, value } = await reader.read()
      if (done) { return }
      let chunk = decoder.decode(value)

      if (pending !== undefined) {
         chunk = pending + chunk
      }

      let lines = chunk.split("\n")
      lines = lines.filter( (line) => line !== "" )

      if (lines && lines[-1] && chunk && lines[-1][-1] === chunk[-1]) {
         pending = lines.pop()
      } else {
         pending = undefined
      }

      for (let i = 0; i < lines.length; i++) { yield lines[i] }

   }

}

export const ChatStore = writable<Chat[]>()
export const UserStore = writable<User>()

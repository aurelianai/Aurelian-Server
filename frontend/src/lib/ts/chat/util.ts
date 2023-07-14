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

export const complete = async (chatid: number): Promise<Message> => {
   const res = await fetch(`/api/chat/${chatid}/complete`, {
      headers: new Headers({ 'content-type': 'application/json' }),
      method: "POST"
   })
   return res.json()
}

export const ChatStore = writable<Chat[]>()
export const UserStore = writable<User>()
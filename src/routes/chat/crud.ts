import type { Chat, Delete, Message, TextGenModel } from '$lib/types'
import { get } from 'svelte/store'
import { selected_session_id } from '$lib/stores'

// Chat
export const list_chats = async (): Promise<Chat[]> => {
   const res = await fetch("/api/chat", { method: "GET" })
   return res.json()
}

export const new_chat = async (): Promise<Chat> => {
   const res = await fetch("/api/chat", { method: "POST" })
   return res.json()
}

export const update_chat = async (id: number, new_name: string) => {
   const body = { id: id, name: new_name }
   const res = await fetch("/api/chat", {
      method: "PATCH",
      body: JSON.stringify(body)
   })
   return res.json()
}

export const delete_chat = async (id: number) => {
   const body: Delete = { id: id }
   await fetch("/api/chat", {
      method: "DELETE",
      body: JSON.stringify(body)
   })
}


// Messages
export const list_messages = async (): Promise<Message[]> => {
   await ssid()
   const res = await fetch(`/api/chat/${get(selected_session_id)}`)
   return res.json()
}

export const new_message = async (role: "USER" | "MODEL", content: string): Promise<Message> => {
   await ssid()
   return { role, content }
}

export const complete = async (): Promise<Message> => {
   await ssid()
   return {
      role: "MODEL",
      content: "Hey, how's it going?"
   }
}


// Models
export const list_models = async (): Promise<TextGenModel[]> => {
   const res = await fetch('/api/models')
   return res.json()
}


// Light Utilities to Wait for State to be Defined, should almost never hit
export const ssid = async () => {
   while (get(selected_session_id) === undefined) {
      await new Promise(r => setTimeout(r, 100))
   }
}
import type { Chat, Delete, Message, TextGenModel } from '$lib/types'

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
   return []
}

export const new_message = async (role: "USER" | "MODEL", content: string): Promise<Message> => {
   return { role, content }
}

export const complete = async (id: number): Promise<Message> => {
   console.log(id)
   return {
      role: "MODEL",
      content: "Hey, how's it going?"
   }
}


// Models
export const list_models = async (): Promise<TextGenModel[]> => {
   return []
}
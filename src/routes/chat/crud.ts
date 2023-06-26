import type { Chat, Delete } from '$lib/types'

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

import type { RequestHandler } from "@sveltejs/kit";
import { redirect, json } from '@sveltejs/kit'

export const GET = (async () => {
   return fetch('google.com')
})
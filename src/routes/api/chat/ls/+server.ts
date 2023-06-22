import type { RequestHandler } from "@sveltejs/kit";

import { redirect, json } from '@sveltejs/kit'

export const POST = (async ({ request }) => {
   return fetch("google.com")
}) satisfies RequestHandler
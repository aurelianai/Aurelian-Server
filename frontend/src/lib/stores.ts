import { writable } from 'svelte/store'

export const selected_session_id = writable<number>();
export const selected_model_id = writable<number>();
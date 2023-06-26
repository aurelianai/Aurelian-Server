<script lang="ts">
	import type { ToastSettings } from '@skeletonlabs/skeleton';
	import { toastStore } from '@skeletonlabs/skeleton';
	import { createEventDispatcher } from 'svelte';

	export let disabled: boolean;

	// Text box size
	let rows: number = 1;
	$: rows = Math.min(countlines(chat_content), 4);

	const countlines = (str: string): number => {
		let num_new_lines = 0;
		for (let i = 0; i < str.length; i++) {
			if (str[i] === '\n') num_new_lines++;
			if (num_new_lines > 4) break;
		}
		return num_new_lines + 1;
	};

	// Chat Send Logic
	let dispatch = createEventDispatcher();
	let chat_content = '';

	let shift_pressed = false;
	let enter_pressed = false;
	let user_send_chat = shift_pressed && enter_pressed;
	$: {
		user_send_chat = shift_pressed && enter_pressed;
		if (user_send_chat) {
			shift_pressed = false;
			enter_pressed = false;
			if (!lint_input()) {
				let t: ToastSettings = {
					message: 'Please write something before sending chat!',
					background: 'variant-filled-error',
					timeout: 750
				};
				toastStore.trigger(t);
			} else {
				dispatch('send_message', {
					message_content: chat_content
				});
				chat_content = '';
			}
		}
	}

	const lint_input = (): boolean => {
		return !(chat_content.length === 0);
	};

	const key_down = (event: KeyboardEvent) => {
		switch (event.key) {
			case 'Shift':
				shift_pressed = true;
				break;
			case 'Enter':
				enter_pressed = true;
				if (shift_pressed) {
					event.preventDefault();
				}
				break;
		}
	};
	const key_up = (event: KeyboardEvent) => {
		switch (event.key) {
			case 'Shift':
				shift_pressed = false;
				break;
			case 'Enter':
				enter_pressed = false;
				break;
		}
	};
</script>

<div class="max-w-2xl px-5 py-2 mx-auto rounded-lg shadow-2xl bg-surface-100-800-token">
	<textarea
		bind:value={chat_content}
		on:keyup={key_up}
		on:keydown={key_down}
		class="p-2 rounded-md textarea"
		{rows}
		placeholder="Type your question here ..."
		{disabled}
	/>
	<p class="my-1 text-xs font-bold text-center">
		<kbd>SHIFT</kbd> + <kbd>ENTER</kbd> to Send Chat
	</p>
</div>

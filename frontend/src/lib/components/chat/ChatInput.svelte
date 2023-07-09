<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { Icon, PaperAirplane } from 'svelte-hero-icons';

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
	let user_send_chat = enter_pressed && !shift_pressed;
	let new_line = shift_pressed && enter_pressed;
	$: {
		user_send_chat = enter_pressed && !shift_pressed;
		new_line = shift_pressed && enter_pressed;
		if (user_send_chat) {
			send_message();
		} else if (new_line) {
			chat_content += '\n';
		}
	}

	const send_message = () => {
		shift_pressed = false;
		enter_pressed = false;
		dispatch('send_message', {
			message_content: chat_content
		});
		chat_content = '';
	};

	const key_down = (event: KeyboardEvent) => {
		switch (event.key) {
			case 'Shift':
				shift_pressed = true;
				break;
			case 'Enter':
				enter_pressed = true;
				event.preventDefault();
				break;
			case 'Tab':
				event.preventDefault();
				chat_content += '\t';
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

<div class="relative flex flex-grow max-w-2xl">
	<textarea
		bind:value={chat_content}
		on:keyup={key_up}
		on:keydown={key_down}
		class="p-4 border-none rounded-md shadow-2xl resize-none textarea bg-surface-100-800-token"
		{rows}
		placeholder="Prompt"
		{disabled}
	/>
	<button
		class="absolute p-2 rounded-md w-9 h-9 btn btn-icon variant-filled-primary bottom-[10px] right-2"
		disabled={chat_content === ''}
		on:click={send_message}
	>
		<Icon src={PaperAirplane} />
	</button>
</div>

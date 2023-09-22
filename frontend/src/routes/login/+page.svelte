<script lang="ts">
	import type { ActionData, PageData } from './$types';
	import { enhance } from '$app/forms';
	import { LightSwitch, focusTrap } from '@skeletonlabs/skeleton';
	import { Icon, LockClosed, UserCircle} from 'svelte-hero-icons'

	export let data: PageData;
	export let form: ActionData

	let username = form?.username ?? decodeURIComponent(data.username || '');
	let password = decodeURIComponent(data.password || '');
</script>

<svelte:head>
	<title>Aurelian — Log In</title>
</svelte:head>

<div
	class="flex items-center justify-center w-full h-full bg-[url('/login-background.svg')] bg-tertiary-50-200-token px-4"
>
	<div class="fixed top-4 left-4">
		<div class="p-4 rounded-md shadow-xl card">
			<LightSwitch rounded="rounded-md" />
		</div>
	</div>

	<div class="flex flex-col flex-grow max-w-sm space-y-5 md:max-w-md">
		{#if data.ref === 'chat'}
			<div class="flex-grow max-w-xs p-2 mx-auto text-white rounded-md card variant-filled-surface">
				Your Session Expired!
			</div>
		{/if}

		{#each form?.errors ?? [] as error}
			{error.field}
			{error.message}
		{/each}

		<div class="p-5 space-y-2 rounded-md shadow-2xl card">
			<header class="flex flex-col items-center card-header">
				<img src="/logo.svg" alt="Aurelian Logo" class="w-10 h-10" />
				<p class="text-3xl font-bold">Welcome!</p>
				<p class="font-bold text-tertiary-600-300-token">Log in to your account</p>
			</header>
			<section class="p-4">
				<form class="flex flex-col items-center space-y-10" method="post" action="?/login" use:enhance use:focusTrap={true}>
					<div class="flex flex-col w-full space-y-2 justify-left">
						<div
							class="input-group input-group-divider grid-cols-[auto_1fr_auto] rounded-md border-none hover:shadow-md"
						>
							<div class="input-group-shim"><Icon src={UserCircle} class="w-5 h-5 text-surface-400" solid/></div>
							<input
								name="username"
								class="p-3 border-none"
								type="text"
								placeholder="Username"
								value={username}
							/>
						</div>

						<div
							class="input-group input-group-divider grid-cols-[auto_1fr_auto] rounded-md border-none hover:shadow-md"
						>
							<div class="input-group-shim"><Icon src={LockClosed} class="w-5 h-5 text-warning-600" solid/></div>
							<input
								name="password"
								class="p-3 border-none"
								type="password"
								placeholder="Password"
								value={password}
							/>
						</div>
					</div>
					<button
						class="mt-10 font-bold rounded-md btn bg-[radial-gradient(ellipse_at_top_left,_var(--tw-gradient-stops))] variant-gradient-secondary-primary w-36"
						type="submit"	
					>
						Log In
					</button>
				</form>
			</section>
		</div>
	</div>
</div>

<script>
	import { Circle3 } from 'svelte-loading-spinners';
	import Typewriter from 'svelte-typewriter';
</script>

<div class="flex flex-col items-center justify-center w-full space-y-3">
	<Circle3
		size="50"
		ballTopLeft="#adafb2"
		ballBottomRight="#adafb2"
		ballBottomLeft="#b83243"
		ballTopRight="#b83243"
	/>
	<Typewriter mode="loop">
		<p class="font-bold">Generating Response ...</p>
	</Typewriter>
</div>

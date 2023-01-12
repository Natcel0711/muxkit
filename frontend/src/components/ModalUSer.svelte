<script lang="ts">
	import { enhance } from '$app/forms';

	// Props
	/** Exposes parent props to this component. */
	export let parent: any;
	// Stores
	import { modalStore } from '@skeletonlabs/skeleton';
	// Form Data
	export let formData = {
		id:$modalStore[0].meta?.id,
        name: $modalStore[0].meta?.name,
		email: $modalStore[0].meta?.email,
		password: $modalStore[0].meta?.password

	};
	// We've created a custom submit function to pass the response and close the modal.
	function onFormSubmit(): void {
		if ($modalStore[0].response) $modalStore[0].response(formData);
        $modalStore[0].meta?.onSubmit(formData)
		modalStore.close();
	}
	// Base Classes
	const cBase = 'space-y-4';
	const cForm = 'border border-surface-500 p-4 space-y-4 rounded-container-token';
</script>

<!-- @component This example creates a simple form modal. -->

<div class="modal-example-form {cBase}">
	<!-- Enable for debugging: -->
	<!-- <pre>{JSON.stringify(formData, null, 2)}</pre> -->
    <h3>{!formData.id? 'Insert':'Update'} User</h3>
	<form class="modal-form {cForm}" id="myForm" action="?/InsertOrUpdate" use:enhance method="POST">
		<label>
			<span>Name</span>
			<input type="text" name="name" class="pl-2 w-2/4" bind:value={formData.name} placeholder="Enter name..." />
		</label>
		<label>
			<span>Email</span>
			<input type="text" name="email" class="pl-2 w-2/4" bind:value={formData.email} placeholder="Enter email..." />
		</label>
		<label>
			<span>Password</span>
			<input type="text" name="password" class="pl-2 w-2/4" bind:value={formData.password} placeholder="Enter password..." />
		</label>
        <input type="text" name="id" hidden bind:value={formData.id}/>
	</form>
	<!-- prettier-ignore -->
	<footer class="modal-footer {parent.regionFooter}">
        <button class="btn {parent.buttonNeutral}" on:click={parent.onClose}>{parent.buttonTextCancel}</button>
        <button type="submit" class="btn {parent.buttonPositive}" form="myForm" on:click={onFormSubmit}>Submit Form</button>
    </footer>
</div>
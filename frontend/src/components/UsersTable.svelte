<script lang="ts">
    import { Modal, Table, tableMapperValues, modalStore, Toast, toastStore } from '@skeletonlabs/skeleton';
    import type { TableSource, ModalSettings, ModalComponent, ToastSettings } from '@skeletonlabs/skeleton';
    import ModalUSer from '../components/ModalUSer.svelte';
    export let users
    const tableSimple: TableSource = {
        // A list of heading labels.
        head: ['ID', 'Name'],
        // The data visibly shown in your table body UI.
        body: tableMapperValues(users, ['id', 'name']),
        // Optional: The data returned when interactive is enabled and a row is clicked.
        meta: tableMapperValues(users, ['id', 'name']),
        // Optional: A list of footer labels.
        //foot: ['Total', '', '<code>31.7747</code>']
    };
const mySelectionHandler = (e) => {
    user.id = e.detail[0]
    user.name = e.detail[1]
    triggerCustomModal()
}
const user = {
    id:0,
    name:""
}


function triggerToast(): void {
	const t: ToastSettings = {
		message: '✅ User updated.',
		// Optional: Presets for primary | secondary | tertiary | warning
		preset: 'primary',
		// Optional: The auto-hide settings
		autohide: true,
		timeout: 5000,
		// Optional: Adds a custom action button
	};
	toastStore.trigger(t);
}
			

function triggerCustomModal(): void {
	const modalComponent: ModalComponent = {
		// Pass a reference to your custom component
		ref: ModalUSer,
		// Add your props as key/value pairs
		props: { background: 'bg-red-500' },
		// Provide default slot content as a template literal
		slot: '<p>Skeleton</p>'
	};
	const d: ModalSettings = {
		type: 'component',
		// NOTE: title, body, response, etc are supported!
		component: modalComponent,
		// Pass abitrary data to the component
		meta: { foo: 'bar', fizz: 'buzz', 
            onSubmit: (data)=>{
                tableSimple.body[data.id - 1][1] = data.name
                triggerToast()
            }, 
        name: user.name, id:user.id 
        }
	};
	modalStore.trigger(d);
}
			
</script>

<Table source={tableSimple} interactive={true} on:selected={mySelectionHandler} />

<Modal />

<Toast />
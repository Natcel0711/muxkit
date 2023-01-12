<script lang="ts">
    import { Modal, Table, tableMapperValues, modalStore, Toast, toastStore } from '@skeletonlabs/skeleton';
    import type { TableSource, ModalSettings, ModalComponent, ToastSettings } from '@skeletonlabs/skeleton';
    import ModalUSer from './ModalUser.svelte';
    export let users
    const tableSimple: TableSource = {
        // A list of heading labels.
        head: ['ID', 'Name', 'Email', 'Password', 'Created At', 'Updated At'],
        // The data visibly shown in your table body UI.
        body: tableMapperValues(users.users, ['id', 'name', 'email', 'password', 'createdat', 'updatedat']),
        // Optional: The data returned when interactive is enabled and a row is clicked.
        meta: tableMapperValues(users.users, ['id', 'name', 'email', 'password', 'createdat', 'updatedat']),
        // Optional: A list of footer labels.
        //foot: ['Total', '', '<code>31.7747</code>']
    };
const mySelectionHandler = (e) => {
    user.id = e.detail[0]
    user.name = e.detail[1]
	user.email = e.detail[2]
	user.password = e.detail[3]
    triggerCustomModal()
}
const user = {
    id:0,
    name:"",
    email:"",
    password:""
}

function UpdateToast(): void {
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
function InsertToast(user): void {
	const t: ToastSettings = {
		message: `✅ ${user} Added.`,
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
                if(data.id){
                    tableSimple.body[data.id - 2][1] = data.name
					tableSimple.body[data.id - 2][2] = data.email
					tableSimple.body[data.id - 2][3] = data.password
                    UpdateToast()
                }else{
					// const newUser = [0, data.name, data.email, data.password, Date.now(), Date.now()]
					// tableSimple.body.push(newUser)
					// tableSimple.meta?.push(newUser)
					// users.users.push({
					// 	id:0,
					// 	name:data.name,
					// 	email:data.email,
					// 	password:data.password,
					// 	createdat:Date.now(),
					// 	updatedat:Date.now()
					// })
                    InsertToast(data.name)
                }
            }, 
        name: user.name, id:user.id, email:user.email, password: user.password 
        }
	};
	modalStore.trigger(d);
}
			
</script>
<button class="btn btn-ringed-primary btn-base ring-2 ring-surface-500 ring-inset m-2" on:click={mySelectionHandler}><strong>ADD</strong>➕</button>
<Table source={tableSimple} interactive={true} on:selected={mySelectionHandler} />

<Modal />

<Toast />
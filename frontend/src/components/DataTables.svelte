<script lang="ts">
    export let data
    import { TabGroup, Tab } from '@skeletonlabs/skeleton';
    const storeTab: Writable<string> = writable('Users');
        import { writable, type Writable } from 'svelte/store';
	import UsersTable from './UsersTable.svelte';
    console.log("Desde datatables:", data)
</script>

<TabGroup selected={storeTab}>
	<Tab value="Users">Users</Tab>
	<Tab value="Products">Products</Tab>
</TabGroup>

{#if $storeTab === 'Users'}<UsersTable users={data.users}/>{/if}
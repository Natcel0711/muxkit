<script lang="ts">
    export let data
    import { TabGroup, Tab } from '@skeletonlabs/skeleton';
    const storeTab: Writable<string> = writable('Users');
        import { writable, type Writable } from 'svelte/store';
	import UsersTable from './UsersTable.svelte';
</script>

<TabGroup selected={storeTab}>
	<Tab value="Users">Users</Tab>
</TabGroup>

{#if $storeTab === 'Users'}<UsersTable users={data.users}/>{/if}
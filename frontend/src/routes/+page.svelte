<script lang="ts">
    export let data
    import { TabGroup, Tab } from '@skeletonlabs/skeleton';
	import JsonToGo from '../components/JsonToGo.svelte';
	import { writable, type Writable } from 'svelte/store';
	import UsersTable from '../components/UsersTable.svelte';
    const storeTab: Writable<string> = writable('a');
    console.log("Del server:", data)
</script>


<TabGroup selected={storeTab}>
	<Tab value="a">JsonToGo</Tab>
	<Tab value="b">Users</Tab>
</TabGroup>

<!-- Conditionally display content -->
{#if $storeTab === 'a'}<JsonToGo/>{/if}
{#if $storeTab === 'b'}<UsersTable users={data.users}/>{/if}
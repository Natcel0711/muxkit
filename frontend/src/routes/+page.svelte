<script lang="ts">
    export let data
    import { TabGroup, Tab } from '@skeletonlabs/skeleton';
	import JsonToGo from '../components/JsonToGo.svelte';
	import { writable, type Writable } from 'svelte/store';
	import DataTables from "../components/DataTables.svelte";
    const storeTab: Writable<string> = writable('a');
</script>


<TabGroup selected={storeTab}>
	<Tab value="a">JsonToGo</Tab>
	<Tab value="b">DataTables</Tab>
</TabGroup>

<!-- Conditionally display content -->
{#if $storeTab === 'a'}<JsonToGo/>{/if}
{#if $storeTab === 'b'}<DataTables data={data} />{/if}
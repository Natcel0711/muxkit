<script lang="ts">
    import { Table, tableMapperValues } from '@skeletonlabs/skeleton';
    import type { TableSource } from '@skeletonlabs/skeleton';
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
let mySelectionHandler
</script>

<Table source={tableSimple} interactive={true} on:selected={mySelectionHandler} />
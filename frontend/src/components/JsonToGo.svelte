<script lang="ts">
import { clipboard } from '@skeletonlabs/skeleton';
import {jsonToGo} from '../utilities/utils'
let robert
let output
let error
function GetJson(data){
    if(!data){
        error = undefined
        return undefined
    }
    let result = jsonToGo(data)
    console.log("Result:",result.go,"Error:", result.error)
    if(result.error){
        error = result.error
        return undefined
    }
    error = ""
    return result.go
}
$: output = GetJson(robert)
</script>
<div>
    <textarea data-clipboard="exampleInput" bind:value={robert} class="bg-gray-200 text-stone-900 rounded-lg shadow-lg p-4 w-2/3 m-4 h-32"/>
</div>
{#if error}
<div class="mx-4">
    <p style="color: red;">*{error}*</p>
</div>
{/if}
<div>
    <textarea disabled data-clipboard="exampleOutput" bind:value={output} class="bg-gray-200 text-stone-900 rounded-lg shadow-lg p-4 w-2/3 m-4 h-32"/>
</div>
<!-- Trigger -->
<div class="mx-4">
    <button class="btn btn-filled-secondary btn-base" use:clipboard={{ input: 'exampleOutput' }}>Copy</button>
</div>

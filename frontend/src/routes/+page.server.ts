export async function load() {
    const response = await fetch('http://localhost:8080/foo', {mode:"cors", credentials:"same-origin"});
    let data = await response.json()
    return {
      posts: data
    };
  }
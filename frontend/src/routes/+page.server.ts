export async function load() {
    const response = await fetch('http://localhost:8080/users', {mode:"cors", credentials:"same-origin"});
    let data = await response.json()
    return {
      users: data.EmployeeList
    };
  }
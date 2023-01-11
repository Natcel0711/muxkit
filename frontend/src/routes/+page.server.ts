import type { User } from '../models/User';

export async function load() {
	const response = await fetch('http://localhost:8080/users', {
		mode: 'cors',
		credentials: 'same-origin'
	});
	let data = await response.json();
  data.EmployeeList.sort(function(a, b) { 
    return a.id - b.id  ||  a.name.localeCompare(b.name);
  });
	return {
		users: data.EmployeeList
	};
}

export const actions = {
	Update: async ({ request }) => {
		//get form data
		const data = await request.formData();
		//build user obj
		const user: User = {
			Id: Number(data.get('id')),
			Name: data.get('name')
		};
		//http request to update
		const response = await fetch('http://localhost:8080/users', {
			method: 'PUT',
			mode: 'cors',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(user)
		});
		const res = await response.json();
	}
};

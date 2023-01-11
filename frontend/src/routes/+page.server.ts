import type { User } from '../models/User';

export async function load() {
	const response = await fetch('http://localhost:8080/users', {
		mode: 'cors',
		credentials: 'same-origin'
	});
	let data = await response.json();
	console.log("Usuarios:",data)
	data.sort(function (a, b) {
		return a.id - b.id || a.name.localeCompare(b.name);
	});
	return {
		users: data
	};
}

export const actions = {
	Update: async (user) => {
		//build user obj
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
	},
	Insert: async (user) => {
		//build user obj
		try {
			console.log("Inserting user:",user, "\n JSON:", JSON.stringify(user))
			//http request to insert
			const response = await fetch('http://localhost:8080/users', {
				method: 'POST',
				mode: 'cors',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(user)
			});
			const res = await response.json();
			console.log(res);
		} catch (error) {
			console.log(error);
		}
	},
	InsertOrUpdate: async ({ request }) => {
		const data = await request.formData();
		//build user obj
		const user: User = {
			Id: Number(data.get('id')),
			Name: data.get('name'),
			Password: data.get('password'),
			Email: data.get('email'),
			CreatedAt: data.get('CreatedAt'),
			UpdatedAt: data.get('UpdatedAt'),
			DeletedAt: data.get('DeletedAt'),
		};

		if (user.Id) actions.Update(user);
		else if (!user.Id) actions.Insert(user);

		return {
			Added: user
		}
	}
};

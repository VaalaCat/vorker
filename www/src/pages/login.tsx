import { HeaderComponent } from '@/components/header';
import { Layout } from '@/components/layout';
import * as React from 'react';



export default function SignIn() {
	const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
		event.preventDefault();
		const data = new FormData(event.currentTarget);
		console.log({
			email: data.get('email'),
			password: data.get('password'),
		});
	};

	return (
		<Layout
			header={<HeaderComponent />}
			side={<></>}
			main={<></>}
		></Layout>
	);
}
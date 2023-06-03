import { RegisterRequest } from "@/types/body";
import { Button, Card, Form, Toast } from "@douyinfe/semi-ui";
import * as api from '@/api/auth';

export function RegisterComponent() {
	const handleSubmit = (values: RegisterRequest) => {
		Toast.info('正在注册，请稍等...');
		api.Reg(values).then(res => {
			if (res.status === 0) {
				Toast.success('注册成功，跳转中...');
				window.location.href = '/admin';
			}
		}).catch(err => {
			Toast.error('注册失败');
			console.log(err);
		})
	};

	return <Card style={{
		position: 'absolute', top: '30%', left: '50%', transform: 'translate(-50%, -50%)',
		width: 400,
	}}
	> <Form layout='vertical'
		onSubmit={values => handleSubmit(values)}
	>
			<Form.Input field='UserName' label='用户名' labelPosition="inset" style={{ width: 200 }} />
			<Form.Input
				field='Password'
				labelPosition="inset"
				label={{ text: '密码' }}
				style={{ width: 200 }}
			/>
			<Form.Input
				field='Email'
				labelPosition="inset"
				label={{ text: '邮箱' }}
				style={{ width: 200 }}
			/>
			<Button htmlType='submit' type="primary">提交</Button>
		</Form>
	</Card>
}
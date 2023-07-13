import { LoginRequest } from '@/types/body'
import { Button, Card, Form, Input, Toast, Tooltip } from '@douyinfe/semi-ui'
import * as api from '@/api/auth'
import { t } from '@/lib/i18n'
import { useMutation } from '@tanstack/react-query'

export function LoginComponent() {
  const handleSubmit = useMutation(async (values: LoginRequest) => {
    return api
      .login(values)
      .then((res) => {
        if (res.status === 0) {
          Toast.success(t.loginSuccess)
          window.location.href = '/admin'
        }
      })
      .catch((err) => {
        Toast.error(t.loginFailed)
        console.log(err)
      })
  })

  return (
    <Card
      style={{
        position: 'absolute',
        top: '30%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        width: 400,
      }}
    >
      {' '}
      <Form
        layout="vertical"
        onSubmit={(values) => handleSubmit.mutate(values)}
      >
        <Form.Input
          field="UserName"
          label={{ text: t.username }}
          labelPosition="inset"
          style={{ width: 200 }}
        />
        <Form.Input
          field="Password"
          labelPosition="inset"
          mode="password"
          label={{ text: t.password }}
          style={{ width: 200 }}
        />
        <Button
          htmlType="submit"
          type="primary"
          loading={handleSubmit.isLoading}
        >
          {t.submit}
        </Button>
      </Form>
    </Card>
  )
}

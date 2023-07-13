import { RegisterRequest } from '@/types/body'
import { Button, Card, Form, Toast } from '@douyinfe/semi-ui'
import * as api from '@/api/auth'
import { t } from '@/lib/i18n'
import { useMutation } from '@tanstack/react-query'

export function RegisterComponent() {
  const handleSubmit = useMutation((values: RegisterRequest) => {
    return api
      .register(values)
      .then((res) => {
        if (res.status === 0) {
          Toast.success(t.registerSuccess)
          window.location.href = '/admin'
        }
      })
      .catch((err) => {
        Toast.error(t.registerSuccess)
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
        <Form.Input
          field="Email"
          labelPosition="inset"
          label={{ text: t.email }}
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

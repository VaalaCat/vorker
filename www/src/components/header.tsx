import { $vorkerSettings } from '@/store/workers'
import { Avatar, Button, ButtonGroup, Nav, Toast } from '@douyinfe/semi-ui'
import { useQuery } from '@tanstack/react-query'
import { useStore } from '@nanostores/react'
import { LuFunctionSquare } from 'react-icons/lu'
import * as api from '@/api/workers'
import * as auth from '@/api/auth'
import { useEffect } from 'react'
import { $user } from '@/store/userState'
import { useRouter } from 'next/router'
import { CH } from '@/lib/color'
import { IconMenu } from '@douyinfe/semi-icons'
import { $expandSidebar } from './sidebar'
import { t } from '@/lib/i18n'

export const HeaderComponent = () => {
  const user = useStore($user)
  const router = useRouter()

  const { data: appconf } = useQuery(['getAppConf'], () => {
    return api.getAppConfig()
  })
  const { data: userinfo } = useQuery(['getUserInfo'], () => {
    return auth.getUserInfo()
  })

  useEffect(() => {
    $vorkerSettings.set(appconf)
  }, [appconf])

  useEffect(() => {
    if (userinfo) {
      $user.set(userinfo)
    }
  }, [userinfo])

  useEffect(() => {
    if (router.asPath !== '/login' && !user) {
      Toast.warning(t.notLoggedInPrompt)
      router.push({
        pathname: '/login',
      })
    }
  }, [router, user])

  return (
    <Nav mode="horizontal" defaultSelectedKeys={['Home']}>
      <Nav.Header>
        <LuFunctionSquare color="#7f7f7f" style={{ fontSize: 36 }} />
        <span
          className="text-xl ml-2"
          style={{ fontFamily: 'trebuchet ms' }}
          onClick={() => {
            router.push('/admin')
          }}
        >
          Vorker
        </span>
      </Nav.Header>
      <Nav.Footer>
        <ButtonGroup aria-label="header button">
          {!userinfo && (
            <Button
              type="primary"
              theme="borderless"
              onClick={() => {
                router.push({ pathname: '/login' })
              }}
              className="pointer-events-auto"
            >
              {t.login}
            </Button>
          )}
          {!userinfo && appconf?.EnableRegister && (
            <Button
              type="primary"
              theme="borderless"
              onClick={() => {
                router.push({ pathname: '/register' })
              }}
              className="pointer-events-auto"
            >
              {t.register}
            </Button>
          )}
          {userinfo && (
            <Avatar
              size="small"
              shape="square"
              style={{ background: CH.hex(JSON.stringify(userinfo)) }}
            >
              {user?.userName?.slice(0, 2).toUpperCase()}
            </Avatar>
          )}
          {userinfo && (
            <Button
              type="primary"
              theme="borderless"
              onClick={() => {
                auth.logout()
                window.location.reload()
              }}
              className="pointer-events-auto"
            >
              {t.logout}
            </Button>
          )}
          <div className="md:hidden">
            <Button
              theme="borderless"
              icon={<IconMenu />}
              onClick={() => {
                $expandSidebar.set(!$expandSidebar.get())
              }}
            />
          </div>
        </ButtonGroup>
      </Nav.Footer>
    </Nav>
  )
}

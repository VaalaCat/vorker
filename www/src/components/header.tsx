import { VorkerSettingsAtom } from "@/store/workers"
import { Button, ButtonGroup, Nav } from "@douyinfe/semi-ui"
import { useQuery } from "@tanstack/react-query"
import { useAtom } from "jotai"
import { LuFunctionSquare } from "react-icons/lu"
import * as api from "@/api/workers"
import { useEffect } from "react"

export const HeaderComponent = () => {
	const [appConf, setAppConf] = useAtom(VorkerSettingsAtom)
	const { data: appconf } = useQuery(['getAppConf'], () => {
		return api.GetAppConfig()
	})

	useEffect(() => {
		setAppConf(appconf)
	}, [appconf])

	return (
		<Nav mode="horizontal" defaultSelectedKeys={['Home']} >
			<Nav.Header>
				<LuFunctionSquare color="#7f7f7f" style={{ fontSize: 36 }} />
				<span className="text-xl ml-2" style={{ fontFamily: 'trebuchet ms' }}>
					Vorker
				</span>
			</Nav.Header>
			<Nav.Footer>
				<ButtonGroup aria-label="header button">
					<Button type="primary" theme="borderless">
						登陆
					</Button>
				</ButtonGroup>
			</Nav.Footer>
		</Nav>
	)
}
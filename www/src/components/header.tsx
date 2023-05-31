import { Button, ButtonGroup, Nav } from "@douyinfe/semi-ui"
import { LuFunctionSquare } from "react-icons/lu"

export const HeaderComponent = () => {
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
						Admin
					</Button>
					<Button type="primary" theme="borderless">
						Login
					</Button>
				</ButtonGroup>
			</Nav.Footer>
		</Nav>
	)
}
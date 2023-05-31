import { Nav } from "@douyinfe/semi-ui"
import { IconWrench, IconPercentage } from '@douyinfe/semi-icons';

export const SideBarComponent = () => {
	return (
		<Nav
			style={{ height: '93vh' }}
			items={[
				{ itemKey: 'workers', text: 'Workers', icon: <IconPercentage /> },
				{ itemKey: 'settings', text: '设置', icon: <IconWrench /> },
			]}
			onSelect={(data) => console.log('trigger onSelect: ', data)}
			onClick={(data) => console.log('trigger onClick: ', data)}
			footer={{ collapseButton: true }}
			defaultIsCollapsed={true}
			defaultSelectedKeys={['workers']}
		/>
	)
}
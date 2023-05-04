import { MonacoEditor } from "@/components/editor";
import { Layout } from "@/components/layout";
import { Avatar, Box, Button, ButtonGroup, IconButton, Stack } from "@mui/material";
import EngineeringIcon from '@mui/icons-material/Engineering';
import SettingsIcon from '@mui/icons-material/Settings';

const sideButtons = [
	<Button key="Workers" startIcon={<EngineeringIcon />}>Workers</Button>,
	<Button key="Settings" startIcon={<SettingsIcon />} >Settings</ Button>,
];

const SideBarComponent = () => {
	return <Stack marginTop={2}>
		<Box margin="auto" marginBottom={2}>
			<Avatar alt="Vorker" src="https://oss.vaala.tech/vaalacat/oss/img/avatar.png" />
		</Box>
		<ButtonGroup fullWidth
			variant="text" orientation="vertical"
			size="small" aria-label="admin sidebar">
			{sideButtons}
		</ButtonGroup>
	</Stack>
}

const HeaderComponent = () => {
	return <ButtonGroup variant="outlined" aria-label="header button">
		<Button >Home</Button>
		<Button >Admin</ Button>
	</ButtonGroup>
}

export default function Admin() {
	return (
		<Layout header={<HeaderComponent />} side={<SideBarComponent />} main={<MonacoEditor />} />
	)
}
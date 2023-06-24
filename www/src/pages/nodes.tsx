import { HeaderComponent } from "@/components/header";
import { Layout } from "@/components/layout";
import { NodesComponent } from "@/components/nodes";
import { SideBarComponent } from "@/components/sidebar";

export default function Register() {
	return (
		<Layout
			header={<HeaderComponent />}
			side={<SideBarComponent selected="status" />}
			main={<NodesComponent />}
		></Layout>
	)
}

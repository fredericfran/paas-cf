package scripts_test

import (
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("ExtractTerraformStateToYAML", func() {

	var (
		cmdInput string
		session  *gexec.Session
	)

	JustBeforeEach(func() {
		command := exec.Command("./extract_terraform_state_to_yaml.rb")
		command.Stdin = strings.NewReader(cmdInput)

		var err error
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("with a valid tfstate file", func() {
		BeforeEach(func() {
			cmdInput = `
{
    "version": 3,
    "terraform_version": "0.7.3",
    "serial": 5,
    "lineage": "bfa4e77c-4e4e-462e-92a9-dba07be0f409",
    "modules": [
        {
            "path": [
                "root"
            ],
            "outputs": {
                "foo_dns_name": {
                    "sensitive": false,
                    "type": "string",
                    "value": "foo.example.com"
                },
                "foo_elastic_ip": {
                    "sensitive": false,
                    "type": "string",
                    "value": "10.210.221.248"
                },
                "foo_security_group_id": {
                    "sensitive": false,
                    "type": "string",
                    "value": "sg-12345678"
                }
            },
			"resources": {},
			"depends_on": []
		}
	]
}
			`
		})

		It("returns a YAML containing all the output values", func() {
			Eventually(session).Should(gexec.Exit(0))
			Expect(session.Out).To(gbytes.Say("terraform_outputs_foo_dns_name: foo.example.com"))
			Expect(session.Out).To(gbytes.Say("terraform_outputs_foo_elastic_ip: 10.210.221.248"))
			Expect(session.Out).To(gbytes.Say("terraform_outputs_foo_security_group_id: sg-12345678"))
		})
	})

	Context("with an old-format tfstate file", func() {
		BeforeEach(func() {
			cmdInput = `
{
    "version": 1,
    "serial": 5,
    "modules": [
        {
            "path": [
                "root"
            ],
            "outputs": {
                "foo_dns_name": "foo.example.com",
                "foo_elastic_ip": "10.210.221.248",
                "foo_security_group_id": "sg-12345678"
            },
			"resources": {},
			"depends_on": []
		}
	]
}
			`
		})

		It("Exits non-zero with an error on STDERR", func() {
			Eventually(session).Should(gexec.Exit())
			Expect(session.ExitCode()).NotTo(Equal(0))
			Expect(session.Out.Contents()).To(BeEmpty())
		})
	})
})

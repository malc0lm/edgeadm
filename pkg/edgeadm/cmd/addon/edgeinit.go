package addon

import (
	"github.com/spf13/cobra"
	"github.com/superedge/edgeadm/pkg/edgeadm/cmd"
	"k8s.io/klog/v2"

	"github.com/superedge/edgeadm/pkg/edgeadm/common"
	"github.com/superedge/edgeadm/pkg/edgeadm/constant"
	"github.com/superedge/edgeadm/pkg/util"
)

func NewEdgeInitCMD() *cobra.Command {
	action := addonAction{
		edgeadmConfig: &cmd.EdgeadmConfig{},
	}
	cmd := &cobra.Command{
		Use:   "edgeinit",
		Short: "Init Kubernetes cluster for join superedge node",
		Run: func(cmd *cobra.Command, args []string) {
			if err := action.complete(); err != nil {
				util.OutPutMessage(err.Error())
				return
			}

			if err := action.runInit(); err != nil {
				util.OutPutMessage(err.Error())
				return
			}
		},
	}
	action.flags = cmd.Flags()
	cmd.Flags().StringVar(&action.manifestDir, "manifest-dir", "",
		"Manifests document of edge kubernetes cluster.")

	cmd.Flags().StringVar(&action.caCertFile, "ca.cert", constant.KubeadmCertPath,
		"The root certificate file for cluster.")

	cmd.Flags().StringVar(&action.caKeyFile, "ca.key", constant.KubeadmKeyPath,
		"The root certificate key file for cluster.")
	cmd.Flags().StringVar(&action.masterPublicAddr, "master-public-addr", "",
		"The public IP for control plane")
	cmd.Flags().StringArrayVar(&action.certSANs, "certSANs", []string{""},
		"The cert SAN")
	cmd.Flags().StringVar(
		&action.edgeadmConfig.Version, constant.EdgeVersion, constant.Version, "Superedge realted images' version.",
	)

	cmd.Flags().StringVar(
		&action.edgeadmConfig.EdgeImageRepository, constant.EdgeImageRepository, constant.ImageRepository, "Superedge related images registry, seperated from the default --image-repository (k8s.gcr.io).",
	)

	cmd.Flags().StringVar(
		&action.edgeadmConfig.EdgeVirtualAddr, constant.EdgeVirtualAddr, constant.DefaultEdgeVirtualAddr, "Superedge related images registry, seperated from the default --image-repository (k8s.gcr.io).",
	)
	return cmd
}

func NewEdgeRestoreCMD() *cobra.Command {
	action := addonAction{}
	cmd := &cobra.Command{
		Use:   "edgeinit",
		Short: "Restore Kubernetes cluster for join superedge node",
		Run: func(cmd *cobra.Command, args []string) {
			if err := action.complete(); err != nil {
				util.OutPutMessage(err.Error())
				return
			}

			if err := action.runRestore(); err != nil {
				util.OutPutMessage(err.Error())
				return
			}
		},
	}
	action.flags = cmd.Flags()
	cmd.Flags().StringVar(&action.manifestDir, "manifest-dir", "",
		"Manifests document of edge kubernetes cluster.")

	cmd.Flags().StringVar(&action.caCertFile, "ca.cert", constant.KubeadmCertPath,
		"The root certificate file for cluster. (default \"/etc/kubernetes/pki/ca.crt\")")

	cmd.Flags().StringVar(&action.caKeyFile, "ca.key", constant.KubeadmKeyPath,
		"The root certificate key file for cluster. (default \"/etc/kubernetes/pki/ca.key\")")
	cmd.Flags().StringVar(&action.masterPublicAddr, "master-public-addr", "",
		"The public IP for control plane")
	cmd.Flags().StringArrayVar(&action.certSANs, "certSANs", []string{""},
		"The cert SAN")

	return cmd
}

func (a *addonAction) runInit() error {
	klog.Info("Start init Kubernetes cluster for join superedge node")
	return common.EdgeClusterInit(a.clientSet, a.manifestDir, a.caCertFile, a.caKeyFile, a.masterPublicAddr, a.certSANs, a.kubeConfig, a.edgeadmConfig)
}

func (a *addonAction) runRestore() error {
	klog.Info("Start restore Kubernetes cluster for join superedge node")
	return common.EdgeClusterRestore(a.clientSet, a.manifestDir, a.caCertFile, a.caKeyFile, a.masterPublicAddr, a.certSANs, a.kubeConfig)
}

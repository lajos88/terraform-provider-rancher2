// Copyright 2020 Oracle and/or its affiliates.

package rancher2

// Flatteners

func flattenClusterOKEConfig(in *OracleKubernetesEngineConfig, p []interface{}) ([]interface{}, error) {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}, nil
	}

	if len(in.CompartmentID) > 0 {
		obj["compartment_id"] = in.CompartmentID
	}

	if len(in.Description) > 0 {
		obj["description"] = in.Description
	}

	obj["enable_kubernetes_dashboard"] = in.EnableKubernetesDashboard
	obj["enable_private_nodes"] = in.PrivateNodes

	if len(in.Fingerprint) > 0 {
		obj["fingerprint"] = in.Fingerprint
	}

	if len(in.KubernetesVersion) > 0 {
		obj["kubernetes_version"] = in.KubernetesVersion
	}

	if len(in.ServiceLBSubnet1Name) > 0 {
		obj["load_balancer_subnet_name_1"] = in.ServiceLBSubnet1Name
	}

	if len(in.ServiceLBSubnet2Name) > 0 {
		obj["load_balancer_subnet_name_2"] = in.ServiceLBSubnet2Name
	}

	if len(in.NodeImage) > 0 {
		obj["node_image"] = in.NodeImage
	}

	if len(in.NodePoolSubnetDNSDomainName) > 0 {
		obj["node_pool_dns_domain_name"] = in.NodePoolSubnetDNSDomainName
	}

	if len(in.NodePoolSubnetName) > 0 {
		obj["node_pool_subnet_name"] = in.NodePoolSubnetName
	}

	if len(in.NodePublicSSHKeyContents) > 0 {
		obj["node_public_key_contents"] = in.NodePublicSSHKeyContents
	}

	if len(in.NodeShape) > 0 {
		obj["node_shape"] = in.NodeShape
	}

	if len(in.PrivateKeyContents) > 0 {
		obj["private_key_contents"] = in.PrivateKeyContents
	}

	if len(in.PrivateKeyPassphrase) > 0 {
		obj["private_key_passphrase"] = in.PrivateKeyPassphrase
	}

	if in.QuantityOfSubnets > 0 {
		obj["quantity_of_node_subnets"] = int(in.QuantityOfSubnets)
	}

	if in.QuantityPerSubnet > 0 {
		obj["quantity_per_subnet"] = int(in.QuantityPerSubnet)
	}

	if len(in.Region) > 0 {
		obj["region"] = in.Region
	}

	if len(in.ServiceSubnetDNSDomainName) > 0 {
		obj["service_dns_domain_name"] = in.ServiceSubnetDNSDomainName
	}

	obj["skip_vcn_delete"] = in.SkipVCNDelete

	if len(in.TenancyID) > 0 {
		obj["tenancy_id"] = in.TenancyID
	}

	if len(in.UserOCID) > 0 {
		obj["user_ocid"] = in.UserOCID
	}

	if len(in.VCNName) > 0 {
		obj["vcn_name"] = in.VCNName
	}

	if len(in.WorkerNodeIngressCidr) > 0 {
		obj["worker_node_ingress_cidr"] = in.WorkerNodeIngressCidr
	}

	return []interface{}{obj}, nil
}

// Expanders

func expandClusterOKEConfig(p []interface{}, name string) (*OracleKubernetesEngineConfig, error) {
	obj := &OracleKubernetesEngineConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := p[0].(map[string]interface{})

	obj.DisplayName = name
	obj.Name = name
	obj.DriverName = clusterDriverOKE

	if v, ok := in["compartment_id"].(string); ok && len(v) > 0 {
		obj.CompartmentID = v
	}

	if v, ok := in["description"].(string); ok && len(v) > 0 {
		obj.Description = v
	}

	if v, ok := in["enable_kubernetes_dashboard"].(bool); ok {
		obj.EnableKubernetesDashboard = v
	}

	if v, ok := in["enable_private_nodes"].(bool); ok {
		obj.PrivateNodes = v
	}

	if v, ok := in["fingerprint"].(string); ok && len(v) > 0 {
		obj.Fingerprint = v
	}

	if v, ok := in["kubernetes_version"].(string); ok && len(v) > 0 {
		obj.KubernetesVersion = v
	}

	if v, ok := in["load_balancer_subnet_name_1"].(string); ok && len(v) > 0 {
		obj.ServiceLBSubnet1Name = v
	}

	if v, ok := in["load_balancer_subnet_name_2"].(string); ok && len(v) > 0 {
		obj.ServiceLBSubnet2Name = v
	}

	if v, ok := in["name"].(string); ok && len(v) > 0 {
		obj.Name = v
	}

	if v, ok := in["node_image"].(string); ok && len(v) > 0 {
		obj.NodeImage = v
	}

	if v, ok := in["node_pool_dns_domain_name"].(string); ok && len(v) > 0 {
		obj.NodePoolSubnetDNSDomainName = v
	}

	if v, ok := in["node_pool_subnet_name"].(string); ok && len(v) > 0 {
		obj.NodePoolSubnetName = v
	}

	if v, ok := in["node_public_key_contents"].(string); ok && len(v) > 0 {
		obj.NodePublicSSHKeyContents = v
	}

	if v, ok := in["node_shape"].(string); ok && len(v) > 0 {
		obj.NodeShape = v
	}

	if v, ok := in["private_key_contents"].(string); ok && len(v) > 0 {
		obj.PrivateKeyContents = v
	}

	if v, ok := in["private_key_passphrase"].(string); ok && len(v) > 0 {
		obj.PrivateKeyPassphrase = v
	}

	if v, ok := in["quantity_of_node_subnets"].(int); ok && v > 0 {
		obj.QuantityOfSubnets = int64(v)
	}

	if v, ok := in["quantity_per_subnet"].(int); ok && v > 0 {
		obj.QuantityPerSubnet = int64(v)
	}

	if v, ok := in["region"].(string); ok && len(v) > 0 {
		obj.Region = v
	}

	if v, ok := in["service_dns_domain_name"].(string); ok && len(v) > 0 {
		obj.ServiceSubnetDNSDomainName = v
	}

	if v, ok := in["skip_vcn_delete"].(bool); ok {
		obj.SkipVCNDelete = v
	}

	if v, ok := in["tenancy_id"].(string); ok && len(v) > 0 {
		obj.TenancyID = v
	}

	if v, ok := in["user_ocid"].(string); ok && len(v) > 0 {
		obj.UserOCID = v
	}

	if v, ok := in["vcn_name"].(string); ok && len(v) > 0 {
		obj.VCNName = v
	}

	if v, ok := in["worker_node_ingress_cidr"].(string); ok && len(v) > 0 {
		obj.WorkerNodeIngressCidr = v
	}

	return obj, nil
}

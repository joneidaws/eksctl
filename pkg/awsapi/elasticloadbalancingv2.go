// Code generated by ifacemaker; DO NOT EDIT.

package awsapi

import (
	"context"

	. "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

// ELBV2 provides an interface to the AWS ELBV2 service.
type ELBV2 interface {
	// Adds the specified SSL server certificate to the certificate list for the
	// specified HTTPS or TLS listener. If the certificate in already in the
	// certificate list, the call is successful but the certificate is not added again.
	// For more information, see HTTPS listeners
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html)
	// in the Application Load Balancers Guide or TLS listeners
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html)
	// in the Network Load Balancers Guide.
	AddListenerCertificates(ctx context.Context, params *AddListenerCertificatesInput, optFns ...func(*Options)) (*AddListenerCertificatesOutput, error)
	// Adds the specified tags to the specified Elastic Load Balancing resource. You
	// can tag your Application Load Balancers, Network Load Balancers, Gateway Load
	// Balancers, target groups, listeners, and rules. Each tag consists of a key and
	// an optional value. If a resource already has a tag with the same key, AddTags
	// updates its value.
	AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error)
	// Creates a listener for the specified Application Load Balancer, Network Load
	// Balancer, or Gateway Load Balancer. For more information, see the following:
	//
	// *
	// Listeners for your Application Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html)
	//
	// *
	// Listeners for your Network Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html)
	//
	// *
	// Listeners for your Gateway Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-listeners.html)
	//
	// This
	// operation is idempotent, which means that it completes at most one time. If you
	// attempt to create multiple listeners with the same settings, each call succeeds.
	CreateListener(ctx context.Context, params *CreateListenerInput, optFns ...func(*Options)) (*CreateListenerOutput, error)
	// Creates an Application Load Balancer, Network Load Balancer, or Gateway Load
	// Balancer. For more information, see the following:
	//
	// * Application Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html)
	//
	// *
	// Network Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html)
	//
	// *
	// Gateway Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-load-balancers.html)
	//
	// This
	// operation is idempotent, which means that it completes at most one time. If you
	// attempt to create multiple load balancers with the same settings, each call
	// succeeds.
	CreateLoadBalancer(ctx context.Context, params *CreateLoadBalancerInput, optFns ...func(*Options)) (*CreateLoadBalancerOutput, error)
	// Creates a rule for the specified listener. The listener must be associated with
	// an Application Load Balancer. Each rule consists of a priority, one or more
	// actions, and one or more conditions. Rules are evaluated in priority order, from
	// the lowest value to the highest value. When the conditions for a rule are met,
	// its actions are performed. If the conditions for no rules are met, the actions
	// for the default rule are performed. For more information, see Listener rules
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#listener-rules)
	// in the Application Load Balancers Guide.
	CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error)
	// Creates a target group. For more information, see the following:
	//
	// * Target
	// groups for your Application Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html)
	//
	// *
	// Target groups for your Network Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html)
	//
	// *
	// Target groups for your Gateway Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/target-groups.html)
	//
	// This
	// operation is idempotent, which means that it completes at most one time. If you
	// attempt to create multiple target groups with the same settings, each call
	// succeeds.
	CreateTargetGroup(ctx context.Context, params *CreateTargetGroupInput, optFns ...func(*Options)) (*CreateTargetGroupOutput, error)
	// Deletes the specified listener. Alternatively, your listener is deleted when you
	// delete the load balancer to which it is attached.
	DeleteListener(ctx context.Context, params *DeleteListenerInput, optFns ...func(*Options)) (*DeleteListenerOutput, error)
	// Deletes the specified Application Load Balancer, Network Load Balancer, or
	// Gateway Load Balancer. Deleting a load balancer also deletes its listeners. You
	// can't delete a load balancer if deletion protection is enabled. If the load
	// balancer does not exist or has already been deleted, the call succeeds. Deleting
	// a load balancer does not affect its registered targets. For example, your EC2
	// instances continue to run and are still registered to their target groups. If
	// you no longer need these EC2 instances, you can stop or terminate them.
	DeleteLoadBalancer(ctx context.Context, params *DeleteLoadBalancerInput, optFns ...func(*Options)) (*DeleteLoadBalancerOutput, error)
	// Deletes the specified rule. You can't delete the default rule.
	DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error)
	// Deletes the specified target group. You can delete a target group if it is not
	// referenced by any actions. Deleting a target group also deletes any associated
	// health checks. Deleting a target group does not affect its registered targets.
	// For example, any EC2 instances continue to run until you stop or terminate them.
	DeleteTargetGroup(ctx context.Context, params *DeleteTargetGroupInput, optFns ...func(*Options)) (*DeleteTargetGroupOutput, error)
	// Deregisters the specified targets from the specified target group. After the
	// targets are deregistered, they no longer receive traffic from the load balancer.
	DeregisterTargets(ctx context.Context, params *DeregisterTargetsInput, optFns ...func(*Options)) (*DeregisterTargetsOutput, error)
	// Describes the current Elastic Load Balancing resource limits for your Amazon Web
	// Services account. For more information, see the following:
	//
	// * Quotas for your
	// Application Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-limits.html)
	//
	// *
	// Quotas for your Network Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-limits.html)
	//
	// *
	// Quotas for your Gateway Load Balancers
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/quotas-limits.html)
	DescribeAccountLimits(ctx context.Context, params *DescribeAccountLimitsInput, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error)
	// Describes the default certificate and the certificate list for the specified
	// HTTPS or TLS listener. If the default certificate is also in the certificate
	// list, it appears twice in the results (once with IsDefault set to true and once
	// with IsDefault set to false). For more information, see SSL certificates
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#https-listener-certificates)
	// in the Application Load Balancers Guide or Server certificates
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#tls-listener-certificate)
	// in the Network Load Balancers Guide.
	DescribeListenerCertificates(ctx context.Context, params *DescribeListenerCertificatesInput, optFns ...func(*Options)) (*DescribeListenerCertificatesOutput, error)
	// Describes the specified listeners or the listeners for the specified Application
	// Load Balancer, Network Load Balancer, or Gateway Load Balancer. You must specify
	// either a load balancer or one or more listeners.
	DescribeListeners(ctx context.Context, params *DescribeListenersInput, optFns ...func(*Options)) (*DescribeListenersOutput, error)
	// Describes the attributes for the specified Application Load Balancer, Network
	// Load Balancer, or Gateway Load Balancer. For more information, see the
	// following:
	//
	// * Load balancer attributes
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes)
	// in the Application Load Balancers Guide
	//
	// * Load balancer attributes
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#load-balancer-attributes)
	// in the Network Load Balancers Guide
	//
	// * Load balancer attributes
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-load-balancers.html#load-balancer-attributes)
	// in the Gateway Load Balancers Guide
	DescribeLoadBalancerAttributes(ctx context.Context, params *DescribeLoadBalancerAttributesInput, optFns ...func(*Options)) (*DescribeLoadBalancerAttributesOutput, error)
	// Describes the specified load balancers or all of your load balancers.
	DescribeLoadBalancers(ctx context.Context, params *DescribeLoadBalancersInput, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error)
	// Describes the specified rules or the rules for the specified listener. You must
	// specify either a listener or one or more rules.
	DescribeRules(ctx context.Context, params *DescribeRulesInput, optFns ...func(*Options)) (*DescribeRulesOutput, error)
	// Describes the specified policies or all policies used for SSL negotiation. For
	// more information, see Security policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies)
	// in the Application Load Balancers Guide or Security policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies)
	// in the Network Load Balancers Guide.
	DescribeSSLPolicies(ctx context.Context, params *DescribeSSLPoliciesInput, optFns ...func(*Options)) (*DescribeSSLPoliciesOutput, error)
	// Describes the tags for the specified Elastic Load Balancing resources. You can
	// describe the tags for one or more Application Load Balancers, Network Load
	// Balancers, Gateway Load Balancers, target groups, listeners, or rules.
	DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error)
	// Describes the attributes for the specified target group. For more information,
	// see the following:
	//
	// * Target group attributes
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes)
	// in the Application Load Balancers Guide
	//
	// * Target group attributes
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#target-group-attributes)
	// in the Network Load Balancers Guide
	//
	// * Target group attributes
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/target-groups.html#target-group-attributes)
	// in the Gateway Load Balancers Guide
	DescribeTargetGroupAttributes(ctx context.Context, params *DescribeTargetGroupAttributesInput, optFns ...func(*Options)) (*DescribeTargetGroupAttributesOutput, error)
	// Describes the specified target groups or all of your target groups. By default,
	// all target groups are described. Alternatively, you can specify one of the
	// following to filter the results: the ARN of the load balancer, the names of one
	// or more target groups, or the ARNs of one or more target groups.
	DescribeTargetGroups(ctx context.Context, params *DescribeTargetGroupsInput, optFns ...func(*Options)) (*DescribeTargetGroupsOutput, error)
	// Describes the health of the specified targets or all of your targets.
	DescribeTargetHealth(ctx context.Context, params *DescribeTargetHealthInput, optFns ...func(*Options)) (*DescribeTargetHealthOutput, error)
	// Replaces the specified properties of the specified listener. Any properties that
	// you do not specify remain unchanged. Changing the protocol from HTTPS to HTTP,
	// or from TLS to TCP, removes the security policy and default certificate
	// properties. If you change the protocol from HTTP to HTTPS, or from TCP to TLS,
	// you must add the security policy and default certificate properties. To add an
	// item to a list, remove an item from a list, or update an item in a list, you
	// must provide the entire list. For example, to add an action, specify a list with
	// the current actions plus the new action.
	ModifyListener(ctx context.Context, params *ModifyListenerInput, optFns ...func(*Options)) (*ModifyListenerOutput, error)
	// Modifies the specified attributes of the specified Application Load Balancer,
	// Network Load Balancer, or Gateway Load Balancer. If any of the specified
	// attributes can't be modified as requested, the call fails. Any existing
	// attributes that you do not modify retain their current values.
	ModifyLoadBalancerAttributes(ctx context.Context, params *ModifyLoadBalancerAttributesInput, optFns ...func(*Options)) (*ModifyLoadBalancerAttributesOutput, error)
	// Replaces the specified properties of the specified rule. Any properties that you
	// do not specify are unchanged. To add an item to a list, remove an item from a
	// list, or update an item in a list, you must provide the entire list. For
	// example, to add an action, specify a list with the current actions plus the new
	// action.
	ModifyRule(ctx context.Context, params *ModifyRuleInput, optFns ...func(*Options)) (*ModifyRuleOutput, error)
	// Modifies the health checks used when evaluating the health state of the targets
	// in the specified target group. If the protocol of the target group is TCP, TLS,
	// UDP, or TCP_UDP, you can't modify the health check protocol, interval, timeout,
	// or success codes.
	ModifyTargetGroup(ctx context.Context, params *ModifyTargetGroupInput, optFns ...func(*Options)) (*ModifyTargetGroupOutput, error)
	// Modifies the specified attributes of the specified target group.
	ModifyTargetGroupAttributes(ctx context.Context, params *ModifyTargetGroupAttributesInput, optFns ...func(*Options)) (*ModifyTargetGroupAttributesOutput, error)
	// Registers the specified targets with the specified target group. If the target
	// is an EC2 instance, it must be in the running state when you register it. By
	// default, the load balancer routes requests to registered targets using the
	// protocol and port for the target group. Alternatively, you can override the port
	// for a target when you register it. You can register each EC2 instance or IP
	// address with the same target group multiple times using different ports. With a
	// Network Load Balancer, you cannot register instances by instance ID if they have
	// the following instance types: C1, CC1, CC2, CG1, CG2, CR1, CS1, G1, G2, HI1,
	// HS1, M1, M2, M3, and T1. You can register instances of these types by IP
	// address.
	RegisterTargets(ctx context.Context, params *RegisterTargetsInput, optFns ...func(*Options)) (*RegisterTargetsOutput, error)
	// Removes the specified certificate from the certificate list for the specified
	// HTTPS or TLS listener.
	RemoveListenerCertificates(ctx context.Context, params *RemoveListenerCertificatesInput, optFns ...func(*Options)) (*RemoveListenerCertificatesOutput, error)
	// Removes the specified tags from the specified Elastic Load Balancing resources.
	// You can remove the tags for one or more Application Load Balancers, Network Load
	// Balancers, Gateway Load Balancers, target groups, listeners, or rules.
	RemoveTags(ctx context.Context, params *RemoveTagsInput, optFns ...func(*Options)) (*RemoveTagsOutput, error)
	// Sets the type of IP addresses used by the subnets of the specified load
	// balancer.
	SetIpAddressType(ctx context.Context, params *SetIpAddressTypeInput, optFns ...func(*Options)) (*SetIpAddressTypeOutput, error)
	// Sets the priorities of the specified rules. You can reorder the rules as long as
	// there are no priority conflicts in the new order. Any existing rules that you do
	// not specify retain their current priority.
	SetRulePriorities(ctx context.Context, params *SetRulePrioritiesInput, optFns ...func(*Options)) (*SetRulePrioritiesOutput, error)
	// Associates the specified security groups with the specified Application Load
	// Balancer. The specified security groups override the previously associated
	// security groups. You can't specify a security group for a Network Load Balancer
	// or Gateway Load Balancer.
	SetSecurityGroups(ctx context.Context, params *SetSecurityGroupsInput, optFns ...func(*Options)) (*SetSecurityGroupsOutput, error)
	// Enables the Availability Zones for the specified public subnets for the
	// specified Application Load Balancer or Network Load Balancer. The specified
	// subnets replace the previously enabled subnets. When you specify subnets for a
	// Network Load Balancer, you must include all subnets that were enabled
	// previously, with their existing configurations, plus any additional subnets.
	SetSubnets(ctx context.Context, params *SetSubnetsInput, optFns ...func(*Options)) (*SetSubnetsOutput, error)
}


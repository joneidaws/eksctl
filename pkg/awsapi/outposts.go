// Code generated by ifacemaker; DO NOT EDIT.

package awsapi

import (
	"context"

	. "github.com/aws/aws-sdk-go-v2/service/outposts"
)

// Outposts provides an interface to the AWS Outposts service.
type Outposts interface {
	// Cancels the specified order for an Outpost.
	CancelOrder(ctx context.Context, params *CancelOrderInput, optFns ...func(*Options)) (*CancelOrderOutput, error)
	// Creates an order for an Outpost.
	CreateOrder(ctx context.Context, params *CreateOrderInput, optFns ...func(*Options)) (*CreateOrderOutput, error)
	// Creates an Outpost. You can specify either an Availability one or an AZ ID.
	CreateOutpost(ctx context.Context, params *CreateOutpostInput, optFns ...func(*Options)) (*CreateOutpostOutput, error)
	// Creates a site for an Outpost.
	CreateSite(ctx context.Context, params *CreateSiteInput, optFns ...func(*Options)) (*CreateSiteOutput, error)
	// Deletes the specified Outpost.
	DeleteOutpost(ctx context.Context, params *DeleteOutpostInput, optFns ...func(*Options)) (*DeleteOutpostOutput, error)
	// Deletes the specified site.
	DeleteSite(ctx context.Context, params *DeleteSiteInput, optFns ...func(*Options)) (*DeleteSiteOutput, error)
	// Gets information about the specified catalog item.
	GetCatalogItem(ctx context.Context, params *GetCatalogItemInput, optFns ...func(*Options)) (*GetCatalogItemOutput, error)
	// Amazon Web Services uses this action to install Outpost servers. Gets
	// information about the specified connection. Use CloudTrail to monitor this
	// action or Amazon Web Services managed policy for Amazon Web Services Outposts to
	// secure it. For more information, see  Amazon Web Services managed policies for
	// Amazon Web Services Outposts
	// (https://docs.aws.amazon.com/outposts/latest/userguide/security-iam-awsmanpol.html)
	// and  Logging Amazon Web Services Outposts API calls with Amazon Web Services
	// CloudTrail
	// (https://docs.aws.amazon.com/outposts/latest/userguide/logging-using-cloudtrail.html)
	// in the Amazon Web Services Outposts User Guide.
	GetConnection(ctx context.Context, params *GetConnectionInput, optFns ...func(*Options)) (*GetConnectionOutput, error)
	// Gets information about the specified order.
	GetOrder(ctx context.Context, params *GetOrderInput, optFns ...func(*Options)) (*GetOrderOutput, error)
	// Gets information about the specified Outpost.
	GetOutpost(ctx context.Context, params *GetOutpostInput, optFns ...func(*Options)) (*GetOutpostOutput, error)
	// Gets the instance types for the specified Outpost.
	GetOutpostInstanceTypes(ctx context.Context, params *GetOutpostInstanceTypesInput, optFns ...func(*Options)) (*GetOutpostInstanceTypesOutput, error)
	// Gets information about the specified Outpost site.
	GetSite(ctx context.Context, params *GetSiteInput, optFns ...func(*Options)) (*GetSiteOutput, error)
	// Gets the site address of the specified site.
	GetSiteAddress(ctx context.Context, params *GetSiteAddressInput, optFns ...func(*Options)) (*GetSiteAddressOutput, error)
	// Lists the hardware assets for the specified Outpost. Use filters to return
	// specific results. If you specify multiple filters, the results include only the
	// resources that match all of the specified filters. For a filter where you can
	// specify multiple values, the results include items that match any of the values
	// that you specify for the filter.
	ListAssets(ctx context.Context, params *ListAssetsInput, optFns ...func(*Options)) (*ListAssetsOutput, error)
	// Lists the items in the catalog. Use filters to return specific results. If you
	// specify multiple filters, the results include only the resources that match all
	// of the specified filters. For a filter where you can specify multiple values,
	// the results include items that match any of the values that you specify for the
	// filter.
	ListCatalogItems(ctx context.Context, params *ListCatalogItemsInput, optFns ...func(*Options)) (*ListCatalogItemsOutput, error)
	// Lists the Outpost orders for your Amazon Web Services account.
	ListOrders(ctx context.Context, params *ListOrdersInput, optFns ...func(*Options)) (*ListOrdersOutput, error)
	// Lists the Outposts for your Amazon Web Services account. Use filters to return
	// specific results. If you specify multiple filters, the results include only the
	// resources that match all of the specified filters. For a filter where you can
	// specify multiple values, the results include items that match any of the values
	// that you specify for the filter.
	ListOutposts(ctx context.Context, params *ListOutpostsInput, optFns ...func(*Options)) (*ListOutpostsOutput, error)
	// Lists the Outpost sites for your Amazon Web Services account. Use filters to
	// return specific results. Use filters to return specific results. If you specify
	// multiple filters, the results include only the resources that match all of the
	// specified filters. For a filter where you can specify multiple values, the
	// results include items that match any of the values that you specify for the
	// filter.
	ListSites(ctx context.Context, params *ListSitesInput, optFns ...func(*Options)) (*ListSitesOutput, error)
	// Lists the tags for the specified resource.
	ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error)
	// Amazon Web Services uses this action to install Outpost servers. Starts the
	// connection required for Outpost server installation. Use CloudTrail to monitor
	// this action or Amazon Web Services managed policy for Amazon Web Services
	// Outposts to secure it. For more information, see  Amazon Web Services managed
	// policies for Amazon Web Services Outposts
	// (https://docs.aws.amazon.com/outposts/latest/userguide/security-iam-awsmanpol.html)
	// and  Logging Amazon Web Services Outposts API calls with Amazon Web Services
	// CloudTrail
	// (https://docs.aws.amazon.com/outposts/latest/userguide/logging-using-cloudtrail.html)
	// in the Amazon Web Services Outposts User Guide.
	StartConnection(ctx context.Context, params *StartConnectionInput, optFns ...func(*Options)) (*StartConnectionOutput, error)
	// Adds tags to the specified resource.
	TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error)
	// Removes tags from the specified resource.
	UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error)
	// Updates an Outpost.
	UpdateOutpost(ctx context.Context, params *UpdateOutpostInput, optFns ...func(*Options)) (*UpdateOutpostOutput, error)
	// Updates the specified site.
	UpdateSite(ctx context.Context, params *UpdateSiteInput, optFns ...func(*Options)) (*UpdateSiteOutput, error)
	// Updates the address of the specified site. You can't update a site address if
	// there is an order in progress. You must wait for the order to complete or cancel
	// the order. You can update the operating address before you place an order at the
	// site, or after all Outposts that belong to the site have been deactivated.
	UpdateSiteAddress(ctx context.Context, params *UpdateSiteAddressInput, optFns ...func(*Options)) (*UpdateSiteAddressOutput, error)
	// Update the physical and logistical details for a rack at a site. For more
	// information about hardware requirements for racks, see Network readiness
	// checklist
	// (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#checklist)
	// in the Amazon Web Services Outposts User Guide. To update a rack at a site with
	// an order of IN_PROGRESS, you must wait for the order to complete or cancel the
	// order.
	UpdateSiteRackPhysicalProperties(ctx context.Context, params *UpdateSiteRackPhysicalPropertiesInput, optFns ...func(*Options)) (*UpdateSiteRackPhysicalPropertiesOutput, error)
}


# SSM CloudTrail Slack Webhook

Using SSM to trigger CloudTrail. CloudTrail writes to CloudWatch which triggers lambda which triggers slacks Webhook :)

So the flow is as follows:

SSM change => CloudTrail => CloudWatch => Lambda => Slack Webhook

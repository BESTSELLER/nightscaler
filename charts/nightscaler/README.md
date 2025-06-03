<!-- AUTO-GENERATED -->


<table>
<tr>
<th>Property</th>
<th>Description</th>
<th>Type</th>
<th>Default</th>
</tr>
<tr>

<td>replicaCount</td>
<td>

</td>
<td>number</td>
<td>

```yaml
1
```

</td>
</tr>
<tr>

<td>pubsub.projectId</td>
<td>

Project ID where the Pub/Sub topic is located  
REQUIRED

</td>
<td>string</td>
<td>

```yaml
""
```

</td>
</tr>
<tr>

<td>pubsub.listen_topic</td>
<td>

Pub/Sub topic to subscribe to  
REQUIRED

</td>
<td>string</td>
<td>

```yaml
""
```

</td>
</tr>
<tr>

<td>pubsub.publish_topic</td>
<td>

Pub/Sub topic to publish to  
REQUIRED

</td>
<td>string</td>
<td>

```yaml
""
```

</td>
</tr>
<tr>

<td>pubsub.clustername</td>
<td>

Name of the cluster it is running in. Used to identify the cluster in the Pub/Sub message.  
REQUIRED

</td>
<td>string</td>
<td>

```yaml
""
```

</td>
</tr>
<tr>

<td>image.pullPolicy</td>
<td>

</td>
<td>string</td>
<td>

```yaml
IfNotPresent
```

</td>
</tr>
<tr>

<td>image.tag</td>
<td>

Overrides the image tag whose default is the chart appVersion.

</td>
<td>string</td>
<td>

```yaml
0.1.0
```

</td>
</tr>
<tr>

<td>serviceAccount.create</td>
<td>

Specifies whether a service account should be created

</td>
<td>bool</td>
<td>

```yaml
true
```

</td>
</tr>
<tr>

<td>serviceAccount.annotations["iam.gke.io/gcp-service-account"]</td>
<td>

</td>
<td>string</td>
<td>

```yaml
nightscaler@your-project-id.iam.gserviceaccount.com
```

</td>
</tr>
<tr>

<td>podAnnotations</td>
<td>

</td>
<td>object</td>
<td>

```yaml
{}
```

</td>
</tr>
<tr>

<td>podSecurityContext</td>
<td>

</td>
<td>object</td>
<td>

```yaml
{}
```

</td>
</tr>
<tr>

<td>podEnvVars.timezone</td>
<td>

</td>
<td>string</td>
<td>

```yaml
UTC
```

</td>
</tr>
<tr>

<td>podEnvVars.debug</td>
<td>

</td>
<td>bool</td>
<td>

```yaml
false
```

</td>
</tr>
<tr>

<td>podEnvVars.jsonLogging</td>
<td>

</td>
<td>bool</td>
<td>

```yaml
true
```

</td>
</tr>
<tr>

<td>securityContext</td>
<td>

</td>
<td>object</td>
<td>

```yaml
{}
```

</td>
</tr>
<tr>

<td>resources.limits.cpu</td>
<td>

</td>
<td>string</td>
<td>

```yaml
100m
```

</td>
</tr>
<tr>

<td>resources.limits.memory</td>
<td>

</td>
<td>string</td>
<td>

```yaml
128Mi
```

</td>
</tr>
<tr>

<td>resources.requests.cpu</td>
<td>

</td>
<td>string</td>
<td>

```yaml
100m
```

</td>
</tr>
<tr>

<td>resources.requests.memory</td>
<td>

</td>
<td>string</td>
<td>

```yaml
128Mi
```

</td>
</tr>
<tr>

<td>autoscaling.enabled</td>
<td>

</td>
<td>bool</td>
<td>

```yaml
false
```

</td>
</tr>
<tr>

<td>autoscaling.minReplicas</td>
<td>

</td>
<td>number</td>
<td>

```yaml
1
```

</td>
</tr>
<tr>

<td>autoscaling.maxReplicas</td>
<td>

</td>
<td>number</td>
<td>

```yaml
100
```

</td>
</tr>
<tr>

<td>autoscaling.targetCPUUtilizationPercentage</td>
<td>

</td>
<td>number</td>
<td>

```yaml
80
```

</td>
</tr>
</table>

<!-- /AUTO-GENERATED -->

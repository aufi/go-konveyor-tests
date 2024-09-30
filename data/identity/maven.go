package identity

import (
	"github.com/konveyor/tackle2-hub/api"
)

var TackleTestappPrivateMaven = api.Identity{
	Kind: "maven",
	Name: "maven-creds",
	Settings: `<?xml version="1.0" encoding="UTF-8"?>

	<settings xmlns="http://maven.apache.org/SETTINGS/1.2.0"
			  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			  xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.2.0 http://maven.apache.org/xsd/settings-1.2.0.xsd">
	
	  <pluginGroups>
	  </pluginGroups>
	
	  <proxies>
	  </proxies>
	
	  <servers>
		<server>
		   <id>tackle-testapp</id>
		   <username>GITHUB_USER</username>
		   <password>GITHUB_TOKEN</password>
		 </server>
	  </servers>
	  <mirrors>
		<mirror>
		  <id>maven-default-http-blocker</id>
		  <mirrorOf>external:http:*</mirrorOf>
		  <name>Pseudo repository to mirror external repositories initially using HTTP.</name>
		  <url>http://0.0.0.0/</url>
		  <blocked>true</blocked>
		</mirror>
	  </mirrors>
	  <profiles>
		<profile>
		  <id>github</id>
		  <repositories>
			<repository>
			  <id>central</id>
			  <url>https://repo1.maven.org/maven2</url>
			</repository>
			<repository>
			  <id>tackle-testapp</id>
			  <url>https://maven.pkg.github.com/konveyor/tackle-testapp</url>
			  <snapshots>
				<enabled>true</enabled>
			  </snapshots>
			</repository>
		  </repositories>
		</profile>
	  </profiles>
	  <activeProfiles>
		<activeProfile>github</activeProfile>
	  </activeProfiles>
	</settings>`,
}

var TackleTestappPublicMaven = api.Identity{
	Kind: "maven",
	Name: "maven-testapp-public-creds",
	Settings: `<?xml version="1.0" encoding="UTF-8"?>

<settings xmlns="http://maven.apache.org/SETTINGS/1.2.0"
          xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
          xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.2.0 http://maven.apache.org/xsd/settings-1.2.0.xsd">

  <pluginGroups>
  </pluginGroups>

  <proxies>
  </proxies>

  <servers>
    <server>
       <id>tackle-testapp-public</id>
       <username>GITHUB_USER</username>
       <password>GITHUB_TOKEN</password>
     </server>
  </servers>
  <mirrors>
    <mirror>
      <id>maven-default-http-blocker</id>
      <mirrorOf>external:http:*</mirrorOf>
      <name>Pseudo repository to mirror external repositories initially using HTTP.</name>
      <url>http://0.0.0.0/</url>
      <blocked>true</blocked>
    </mirror>
  </mirrors>
  <profiles>
    <profile>
      <id>github</id>
      <repositories>
        <repository>
          <id>central</id>
          <url>https://repo1.maven.org/maven2</url>
        </repository>
        <repository>
          <id>tackle-testapp-public</id>
          <url>https://maven.pkg.github.com/konveyor/tackle-testapp-public</url>
          <snapshots>
            <enabled>true</enabled>
          </snapshots>
        </repository>
      </repositories>
    </profile>
  </profiles>
  <activeProfiles>
    <activeProfile>github</activeProfile>
  </activeProfiles>
</settings>`,
}
